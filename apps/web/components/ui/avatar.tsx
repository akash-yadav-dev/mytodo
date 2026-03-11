import { cn } from "@/lib/utils/cn";

type AvatarProps = {
  name: string;
  src?: string | null;
  size?: "sm" | "md" | "lg";
  className?: string;
};

const sizeMap = {
  sm: "h-7 w-7 text-xs",
  md: "h-9 w-9 text-sm",
  lg: "h-12 w-12 text-base",
};

function initials(name: string): string {
  return name
    .split(" ")
    .map((w) => w[0])
    .slice(0, 2)
    .join("")
    .toUpperCase();
}

export function Avatar({ name, src, size = "md", className }: AvatarProps) {
  if (src) {
    return (
      // eslint-disable-next-line @next/next/no-img-element
      <img
        src={src}
        alt={name}
        className={cn(
          "rounded-full object-cover ring-2 ring-border",
          sizeMap[size],
          className
        )}
      />
    );
  }

  return (
    <span
      aria-label={name}
      className={cn(
        "inline-flex items-center justify-center rounded-full bg-accent/20 font-semibold text-accent ring-2 ring-border",
        sizeMap[size],
        className
      )}
    >
      {initials(name)}
    </span>
  );
}