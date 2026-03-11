// WebSocket client — authenticates via first message (not URL params,
// which are logged in server/proxy access logs).
let ws: WebSocket | null = null;

export function connectWebSocket(token: string, url?: string): WebSocket {
  const endpoint = url ?? (process.env.NEXT_PUBLIC_WS_URL ?? "ws://localhost:8080/ws");
  ws = new WebSocket(endpoint);

  ws.addEventListener("open", () => {
    // Send auth token as the first message so it never appears in server logs
    ws?.send(JSON.stringify({ type: "auth", token }));
  });

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
