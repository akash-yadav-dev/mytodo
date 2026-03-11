"use client";

import React from "react";

export function MainLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="min-h-[calc(100vh-4rem)] bg-bg text-fg">
      <main className="mx-auto max-w-6xl px-6 py-10">{children}</main>
    </div>
  );
}
