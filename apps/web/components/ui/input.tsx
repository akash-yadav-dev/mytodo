import { cn } from "@/lib/utils/cn";

type InputProps = React.InputHTMLAttributes<HTMLInputElement>;

export function Input({ className, ...props }: InputProps) {
	return (
		<input
			className={cn(
				"w-full rounded-2xl border border-border bg-transparent px-4 py-3 text-sm text-fg shadow-sm outline-none transition focus:border-accent",
				className
			)}
			{...props}
		/>
	);
}
