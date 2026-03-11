import type { QueryClient } from "@tanstack/react-query";
import type { WsMessage } from "./events";
import { WS_EVENTS } from "./events";

export function setupWsHandlers(ws: WebSocket, queryClient: QueryClient): void {
  ws.addEventListener("message", (event: MessageEvent<string>) => {
    try {
      const msg = JSON.parse(event.data) as WsMessage;
      switch (msg.event) {
        case WS_EVENTS.ISSUE_CREATED:
        case WS_EVENTS.ISSUE_UPDATED:
        case WS_EVENTS.ISSUE_DELETED:
          queryClient.invalidateQueries({ queryKey: ["issues"] });
          break;
        case WS_EVENTS.BOARD_UPDATED:
          queryClient.invalidateQueries({ queryKey: ["boards"] });
          break;
        case WS_EVENTS.ORG_MEMBER_ADDED:
        case WS_EVENTS.ORG_MEMBER_REMOVED:
          queryClient.invalidateQueries({ queryKey: ["organizations"] });
          break;
      }
    } catch {
      // Ignore malformed messages
    }
  });
}
