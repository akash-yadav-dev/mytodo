"use client";

import { AuthGuard } from "@/components/auth/auth-guard";
import { Button } from "@/components/ui/button";
import { ROUTES } from "@/lib/constants/routes";
import Link from "next/link";

const demoBoards = [
  { id: "board-1", name: "Sprint 1", projectName: "MyTodo Platform", columns: 3 },
  { id: "board-2", name: "Backlog", projectName: "API Gateway", columns: 2 },
];

export default function BoardsPage() {
  return (
    <AuthGuard>
      <div className="mx-auto max-w-6xl px-6 py-10">
        <div className="mb-8 flex items-center justify-between">
          <div>
            <h1 className="text-2xl font-semibold">Boards</h1>
            <p className="mt-1 text-sm text-fg/60">
              Visualize your workflow with Kanban boards.
            </p>
          </div>
          <Button>New board</Button>
        </div>

        {demoBoards.length === 0 ? (
          <div className="flex flex-col items-center justify-center rounded-3xl border border-dashed border-border py-24 text-center">
            <span className="mb-3 text-5xl">📋</span>
            <h3 className="mb-2 text-lg font-medium">No boards yet</h3>
            <p className="mb-6 max-w-xs text-sm text-fg/60">
              Create a board to visualize and manage your project workflow.
            </p>
            <Button>Create board</Button>
          </div>
        ) : (
          <div className="grid gap-4 sm:grid-cols-2 lg:grid-cols-3">
            {demoBoards.map((board) => (
              <Link
                key={board.id}
                href={ROUTES.boardDetail(board.id)}
                className="group rounded-3xl border border-border bg-card p-6 shadow-soft transition hover:-translate-y-1 hover:border-accent/40 hover:shadow-lg"
              >
                <span className="mb-3 block text-3xl">📋</span>
                <h3 className="text-lg font-semibold">{board.name}</h3>
                <p className="mt-1 text-xs text-fg/50">{board.projectName}</p>
                <p className="mt-4 text-sm text-fg/60">
                  {board.columns} columns
                </p>
              </Link>
            ))}
          </div>
        )}
      </div>
    </AuthGuard>
  );
}
