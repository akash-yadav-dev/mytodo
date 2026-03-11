// Event type definitions for real-time updates.
export const WS_EVENTS = {
  ISSUE_CREATED: "issue.created",
  ISSUE_UPDATED: "issue.updated",
  ISSUE_DELETED: "issue.deleted",
  BOARD_UPDATED: "board.updated",
  ORG_MEMBER_ADDED: "org.member.added",
  ORG_MEMBER_REMOVED: "org.member.removed",
} as const;

export type WsEventType = (typeof WS_EVENTS)[keyof typeof WS_EVENTS];

export type WsMessage<T = unknown> = {
  event: WsEventType;
  payload: T;
};
