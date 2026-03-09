"use client";

import { useEffect } from "react";
import { useRouter } from "next/navigation";
import { useAuthContext } from "@/components/auth/auth-provider";

export function AuthGuard({ children }: { children: React.ReactNode }) {
  const { status } = useAuthContext();
  const router = useRouter();

  useEffect(() => {
    if (status === "guest") {
      router.replace("/auth/login");
    }
  }, [status, router]);

  if (status === "loading") {
    return (
      <div className="flex min-h-[60vh] items-center justify-center text-sm text-fg/70">
        Checking your session...
      </div>
    );
  }

  if (status === "guest") {
    return null;
  }

  return <>{children}</>;
}
