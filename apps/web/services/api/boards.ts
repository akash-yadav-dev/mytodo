import { apiFetch } from "@/services/api/client";

export type Board = {
  id: string;
  name: string;
  project_id: string;
  created_at: string;
  updated_at: string;
};

export type BoardColumn = {
  id: string;
  board_id: string;
  name: string;
  order: number;
  issue_ids: string[];
};

export async function listBoards(projectId?: string): Promise<Board[]> {
  const qs = projectId ? `?project_id=${projectId}` : "";
  return apiFetch<Board[]>(`/api/v1/boards${qs}`);
}

export async function getBoard(id: string): Promise<Board> {
  return apiFetch<Board>(`/api/v1/boards/${id}`);
}

export async function createBoard(payload: {
  name: string;
  project_id: string;
}): Promise<Board> {
  return apiFetch<Board>("/api/v1/boards", {
    method: "POST",
    body: JSON.stringify(payload),
  });
}

export async function getBoardColumns(boardId: string): Promise<BoardColumn[]> {
  return apiFetch<BoardColumn[]>(`/api/v1/boards/${boardId}/columns`);
}
