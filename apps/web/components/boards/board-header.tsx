import Link from "next/link";
import { Button } from "@/components/ui/button";
import { ROUTES } from "@/lib/constants/routes";
import type { Board } from "@/services/api/boards";

type Props = {
  board: Board;
  onNewIssue?: () => void;
};

export function BoardHeader({ board, onNewIssue }: Props) {
  return (
    <div className="mb-6 flex items-center justify-between">
      <div>
        <nav className="mb-1 flex items-center gap-1.5 text-xs text-fg/50">
          <Link href={ROUTES.boards} className="hover:text-fg">Boards</Link>
          <span>/</span>
          <span className="text-fg">{board.name}</span>
        </nav>
        <h1 className="text-2xl font-semibold">{board.name}</h1>
      </div>
      {onNewIssue && (
        <Button onClick={onNewIssue}>
          + New issue
        </Button>
      )}
    </div>
  );
}
