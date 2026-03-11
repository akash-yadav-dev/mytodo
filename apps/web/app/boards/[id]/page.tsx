"use client";

import { use } from "react";
import Link from "next/link";
import { AuthGuard } from "@/components/auth/auth-guard";
import { Button } from "@/components/ui/button";
import { ROUTES } from "@/lib/constants/routes";

const demoColumns = [
  {
    id: "todo",
    title: "To Do",
    color: "bg-fg/10",
    cards: [
      { id: "c1", title: "Set up CI pipeline" },
      { id: "c2", title: "Write API docs" },
    ],
  },
  {
    id: "in-progress",
    title: "In Progress",
    color: "bg-amber-100 dark:bg-amber-900/20",
    cards: [
      { id: "c3", title: "Build auth flow" },
      { id: "c4", title: "Design dashboard" },
    ],
  },
  {
    id: "done",
    title: "Done",
    color: "bg-green-100 dark:bg-green-900/20",
    cards: [{ id: "c5", title: "Project setup" }],
  },
];

export default function BoardDetailPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = use(params);

  return (
    <AuthGuard>
      <div className="px-6 py-10">
        <nav className="mb-6 flex items-center gap-2 text-sm text-fg/50">
          <Link href={ROUTES.boards} className="transition hover:text-fg">
            Boards
          </Link>
          <span>/</span>
          <span className="text-fg">Board #{id}</span>
        </nav>

        <div className="mb-6 flex items-center justify-between">
          <h1 className="text-2xl font-semibold">Board #{id}</h1>
          <Button>Add issue</Button>
        </div>

        {/* Kanban columns */}
        <div className="flex gap-4 overflow-x-auto pb-4">
          {demoColumns.map((col) => (
            <div
              key={col.id}
              className="w-72 flex-shrink-0 rounded-3xl border border-border bg-card p-4"
            >
              <div className="mb-3 flex items-center justify-between">
                <h3 className="text-sm font-semibold">{col.title}</h3>
                <span className="rounded-full bg-fg/10 px-2 py-0.5 text-xs">
                  {col.cards.length}
                </span>
              </div>

              <div className="space-y-2">
                {col.cards.map((card) => (
                  <div
                    key={card.id}
                    className="cursor-grab rounded-2xl border border-border bg-bg px-4 py-3 text-sm shadow-soft transition hover:border-accent/40 hover:shadow-md active:cursor-grabbing"
                  >
                    {card.title}
                  </div>
                ))}

                <button className="w-full rounded-2xl border border-dashed border-border py-2 text-xs text-fg/40 transition hover:border-accent/40 hover:text-fg/60">
                  + Add card
                </button>
              </div>
            </div>
          ))}

          {/* Add column */}
          <div className="w-72 flex-shrink-0">
            <button className="flex h-16 w-full items-center justify-center rounded-3xl border border-dashed border-border text-sm text-fg/40 transition hover:border-accent/40 hover:text-fg/60">
              + Add column
            </button>
          </div>
        </div>
      </div>
    </AuthGuard>
  );
}
