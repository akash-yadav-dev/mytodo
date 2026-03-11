// WebSocket client stub — will be wired to the backend WS server.
let ws: WebSocket | null = null;

export function connectWebSocket(token: string, url?: string): WebSocket {
  const endpoint = url ?? (process.env.NEXT_PUBLIC_WS_URL ?? "ws://localhost:8080/ws");
  ws = new WebSocket(`${endpoint}?token=${token}`);
  return ws;
}

export function disconnectWebSocket(): void {
  if (ws && ws.readyState === WebSocket.OPEN) {
    ws.close();
    ws = null;
  }
}

export function getWebSocket(): WebSocket | null {
  return ws;
}
