"use client";

import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { useLogin } from "@/hooks/useAuth/useLogin";
import { ROUTES } from "@/lib/constants/routes";

export default function LoginPage() {
  const router = useRouter();
  const { submit, loading, error } = useLogin();
  const [form, setForm] = useState({ email: "", password: "" });

  const handleSubmit = async (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    const ok = await submit(form);
    if (ok) {
      router.push(ROUTES.dashboard);
    }
  };

  return (
    <div className="space-y-6">
      <div className="space-y-2">
        <h1 className="text-2xl font-semibold">Welcome back</h1>
        <p className="text-sm text-fg/70">Sign in to keep building momentum.</p>
      </div>
      <form className="space-y-4" onSubmit={handleSubmit}>
        <Input
          type="email"
          placeholder="Email"
          value={form.email}
          onChange={(event) => setForm({ ...form, email: event.target.value })}
          required
        />
        <Input
          type="password"
          placeholder="Password"
          value={form.password}
          onChange={(event) => setForm({ ...form, password: event.target.value })}
          required
        />
        {error ? <p className="text-sm text-red-500">{error}</p> : null}
        <Button className="w-full" type="submit" disabled={loading}>
          {loading ? "Signing in..." : "Sign in"}
        </Button>
      </form>
      <p className="text-sm text-fg/70">
        No account yet?{" "}
        <Link href={ROUTES.register} className="font-medium text-fg">
          Create one
        </Link>
      </p>
    </div>
  );
}
