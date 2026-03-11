"use client";

import { useState } from "react";
import Link from "next/link";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Badge } from "@/components/ui/badge";
import { Modal } from "@/components/ui/modal";
import { ProjectForm } from "@/components/forms/project-form";
import { useProjects } from "@/hooks/useProjects/useProjects";
import { useCreateProject } from "@/hooks/useProjects/useCreateProject";
import { ROUTES } from "@/lib/constants/routes";

export default function ProjectsPage() {
  return (
    <AuthGuard>
      <ProjectsContent />
    </AuthGuard>
  );
}

function ProjectsContent() {
  const [createOpen, setCreateOpen] = useState(false);
  const [createError, setCreateError] = useState<string | null>(null);

  const { data, isLoading } = useProjects();
  const createProject = useCreateProject();
  const projects = data?.data ?? [];

  const handleCreate = async (payload: { name: string; description?: string; key: string }) => {
    setCreateError(null);
    try {
      await createProject.mutateAsync(payload);
      setCreateOpen(false);
    } catch (err) {
      setCreateError(err instanceof Error ? err.message : "Failed to create project");
    }
  };

  return (
    <div className="mx-auto max-w-6xl px-6 py-10">
      <div className="mb-8 flex items-start justify-between">
        <div>
          <h1 className="text-2xl font-semibold">Projects</h1>
          <p className="mt-1 text-sm text-fg/60">Organise issues by project or product area.</p>
        </div>
        <Button onClick={() => setCreateOpen(true)}>+ New project</Button>
      </div>

      {isLoading ? (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {Array.from({ length: 6 }).map((_, i) => (
            <div key={i} className="h-40 animate-pulse rounded-2xl bg-card" />
          ))}
        </div>
      ) : projects.length === 0 ? (
        <div className="rounded-2xl border border-dashed border-border py-20 text-center">
          <p className="text-sm text-fg/50">No projects yet.</p>
          <Button className="mt-4" onClick={() => setCreateOpen(true)}>Create your first project</Button>
        </div>
      ) : (
        <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
          {projects.map((project) => (
            <Link key={project.id} href={ROUTES.projectDetail(project.id)}>
              <Card className="group h-full cursor-pointer transition hover:border-accent/40 hover:shadow-lg">
                <div className="mb-3 flex items-center gap-2">
                  <span className="flex h-9 w-9 items-center justify-center rounded-lg bg-accent/10 text-sm font-bold text-accent">
                    {project.key}
                  </span>
                  {project.status && <Badge>{project.status}</Badge>}
                </div>
                <h2 className="font-semibold group-hover:text-accent">{project.name}</h2>
                {project.description && (
                  <p className="mt-1 line-clamp-2 text-sm text-fg/60">{project.description}</p>
                )}
                <p className="mt-3 text-xs text-fg/40">
                  Created {new Date(project.created_at).toLocaleDateString()}
                </p>
              </Card>
            </Link>
          ))}
        </div>
      )}

      <Modal open={createOpen} onClose={() => setCreateOpen(false)} title="New project">
        <ProjectForm
          onSubmit={handleCreate}
          onCancel={() => setCreateOpen(false)}
          loading={createProject.isPending}
          error={createError}
        />
      </Modal>
    </div>
  );
}
