"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthContext } from "@/components/auth/auth-provider";
import { ROUTES } from "@/lib/constants/routes";

export function AuthGuard({ children }: { children: React.ReactNode }) {
  const { status } = useAuthContext();
  const router = useRouter();

  useEffect(() => {
    if (status === "guest") {
      router.push(ROUTES.login);
    }
  }, [status, router]);

  if (status === "loading") {
    return (
      <div className="flex min-h-screen items-center justify-center">
        <div className="h-6 w-6 animate-spin rounded-full border-2 border-accent border-t-transparent" />
      </div>
    );
  }

  if (status === "guest") {
    return null;
  }

  return <>{children}</>;
}
