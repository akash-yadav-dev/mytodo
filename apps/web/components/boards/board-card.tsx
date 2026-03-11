import Link from "next/link";
import { Card } from "@/components/ui/card";
import { ROUTES } from "@/lib/constants/routes";
import type { Board } from "@/services/api/boards";

type Props = { board: Board };

export function BoardCard({ board }: Props) {
  return (
    <Link href={ROUTES.boardDetail(board.id)}>
      <Card className="group cursor-pointer transition hover:border-accent/40 hover:shadow-lg">
        <div className="mb-3 flex h-10 w-10 items-center justify-center rounded-xl bg-accentMuted text-accent text-lg font-bold">
          {board.name[0]?.toUpperCase()}
        </div>
        <h3 className="font-semibold group-hover:text-accent">{board.name}</h3>
        <p className="mt-1 text-xs text-fg/50">
          Created {new Date(board.created_at).toLocaleDateString()}
        </p>
      </Card>
    </Link>
  );
}
