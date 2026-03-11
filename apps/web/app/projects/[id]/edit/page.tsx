"use client";

import { use, useState } from "react";
import Link from "next/link";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Card } from "@/components/ui/card";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ROUTES } from "@/lib/constants/routes";

export default function EditProjectPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);

  return (
    <AuthGuard>
      <EditProjectContent id={id} />
    </AuthGuard>
  );
}

function EditProjectContent({ id }: { id: string }) {
  const [form, setForm] = useState({ name: "", description: "" });
  const [saved, setSaved] = useState(false);

  const handleSubmit = (e: React.FormEvent) => {
    e.preventDefault();
    // TODO: wire to PUT /api/v1/projects/:id
    setSaved(true);
    setTimeout(() => setSaved(false), 2000);
  };

  return (
    <div className="mx-auto max-w-2xl px-6 py-10">
      <nav className="mb-6 flex items-center gap-2 text-sm text-fg/50">
        <Link href={ROUTES.projects} className="transition hover:text-fg">
          Projects
        </Link>
        <span>/</span>
        <Link href={ROUTES.projectDetail(id)} className="transition hover:text-fg">
          Project #{id}
        </Link>
        <span>/</span>
        <span className="text-fg">Edit</span>
      </nav>

      <div className="mb-8">
        <h1 className="text-2xl font-semibold">Edit project</h1>
      </div>

      <Card>
        <form onSubmit={handleSubmit} className="space-y-4">
          <div className="space-y-1">
            <label className="text-xs font-medium text-fg/70">
              Project name
            </label>
            <Input
              value={form.name}
              onChange={(e) => setForm({ ...form, name: e.target.value })}
              placeholder="My awesome project"
            />
          </div>
          <div className="space-y-1">
            <label className="text-xs font-medium text-fg/70">Description</label>
            <Input
              value={form.description}
              onChange={(e) =>
                setForm({ ...form, description: e.target.value })
              }
              placeholder="What is this project about?"
            />
          </div>
          <div className="flex gap-2">
            <Button type="submit" className="flex-1">
              {saved ? "Saved ✓" : "Save changes"}
            </Button>
            <Button
              type="button"
              variant="secondary"
              asChild
              href={ROUTES.projectDetail(id)}
            >
              Cancel
            </Button>
          </div>
        </form>
      </Card>
    </div>
  );
}
