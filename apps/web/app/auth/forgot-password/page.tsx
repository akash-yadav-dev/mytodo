"use client";

import { useState } from "react";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { ROUTES } from "@/lib/constants/routes";
import { APP_NAME } from "@/lib/constants/config";

export default function ForgotPasswordPage() {
  const router = useRouter();
  const [email, setEmail] = useState("");
  const [sent, setSent] = useState(false);
  const [loading, setLoading] = useState(false);

  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setLoading(true);
    // TODO: wire to POST /api/v1/auth/forgot-password when backend adds the endpoint
    await new Promise((r) => setTimeout(r, 800));
    setLoading(false);
    setSent(true);
  };

  if (sent) {
    return (
      <div className="space-y-4 text-center">
        <span className="text-5xl">📬</span>
        <h1 className="text-2xl font-semibold">Check your inbox</h1>
        <p className="text-sm text-fg/70">
          If an account exists for{" "}
          <span className="font-medium text-fg">{email}</span>, you&apos;ll
          receive a password reset link shortly.
        </p>
        <Button
          variant="secondary"
          onClick={() => router.push(ROUTES.login)}
          className="w-full"
        >
          Back to sign in
        </Button>
      </div>
    );
  }

  return (
    <div className="space-y-6">
      <div className="space-y-2">
        <h1 className="text-2xl font-semibold">Forgot your password?</h1>
        <p className="text-sm text-fg/70">
          Enter your email and we&apos;ll send a reset link.
        </p>
      </div>
      <form onSubmit={handleSubmit} className="space-y-4">
        <Input
          type="email"
          placeholder="you@example.com"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
          required
        />
        <Button className="w-full" type="submit" disabled={loading}>
          {loading ? "Sending..." : "Send reset link"}
        </Button>
      </form>
      <p className="text-sm text-fg/70">
        Remember your password?{" "}
        <Link href={ROUTES.login} className="font-medium text-fg">
          Sign in
        </Link>
      </p>
    </div>
  );
}
