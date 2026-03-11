"use client";

import { useEffect, useRef, useState } from "react";
import { cn } from "@/lib/utils/cn";

export type DropdownItem = {
  label: string;
  onClick: () => void;
  destructive?: boolean;
};

type DropdownProps = {
  trigger: React.ReactNode;
  items: DropdownItem[];
  align?: "left" | "right";
  className?: string;
};

export function Dropdown({ trigger, items, align = "right", className }: DropdownProps) {
  const [open, setOpen] = useState(false);
  const ref = useRef<HTMLDivElement>(null);

  useEffect(() => {
    const handler = (e: MouseEvent) => {
      if (ref.current && !ref.current.contains(e.target as Node)) {
        setOpen(false);
      }
    };
    document.addEventListener("mousedown", handler);
    return () => document.removeEventListener("mousedown", handler);
  }, []);

  return (
    <div ref={ref} className={cn("relative inline-block", className)}>
      <div onClick={() => setOpen((v) => !v)} className="cursor-pointer">
        {trigger}
      </div>
      {open && (
        <ul
          className={cn(
            "absolute z-50 mt-2 min-w-[10rem] rounded-2xl border border-border bg-card py-1 shadow-soft",
            align === "right" ? "right-0" : "left-0"
          )}
        >
          {items.map((item, i) => (
            <li key={i}>
              <button
                onClick={() => {
                  item.onClick();
                  setOpen(false);
                }}
                className={cn(
                  "w-full px-4 py-2 text-left text-sm transition hover:bg-accentMuted",
                  item.destructive && "text-red-500"
                )}
              >
                {item.label}
              </button>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}