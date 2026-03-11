"use client";

import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useRegister } from "@/hooks/useAuth/useRegister";
import { ROUTES } from "@/lib/constants/routes";

export function RegisterForm() {
  const router = useRouter();
  const { submit, loading, error } = useRegister();
  const [form, setForm] = useState({ name: "", email: "", password: "" });

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    const ok = await submit(form);
    if (ok) router.push(ROUTES.dashboard);
  };

  return (
    <div className="space-y-6">
      <div className="space-y-1">
        <h1 className="text-2xl font-semibold">Start your workspace</h1>
        <p className="text-sm text-fg/60">Create your account and start tracking.</p>
      </div>

      <form className="space-y-4" onSubmit={handleSubmit}>
        <Input
          type="text"
          placeholder="Full name"
          autoComplete="name"
          value={form.name}
          onChange={(e) => setForm({ ...form, name: e.target.value })}
          required
        />
        <Input
          type="email"
          placeholder="Email address"
          autoComplete="email"
          value={form.email}
          onChange={(e) => setForm({ ...form, email: e.target.value })}
          required
        />
        <Input
          type="password"
          placeholder="Password (min 8 characters)"
          autoComplete="new-password"
          minLength={8}
          value={form.password}
          onChange={(e) => setForm({ ...form, password: e.target.value })}
          required
        />

        {error && (
          <p className="rounded-lg border border-red-200 bg-red-50 px-3 py-2 text-sm text-red-600">
            {error}
          </p>
        )}

        <Button className="w-full" type="submit" disabled={loading}>
          {loading ? "Creating account…" : "Create account"}
        </Button>
      </form>

      <p className="text-center text-sm text-fg/60">
        Already have an account?{" "}
        <Link href={ROUTES.login} className="font-medium text-fg hover:text-accent">
          Sign in
        </Link>
      </p>
    </div>
  );
}
