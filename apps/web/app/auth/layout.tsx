import { AuthLayout } from "@/components/layouts/auth-layout";

export default function AuthRootLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  return <AuthLayout>{children}</AuthLayout>;
}
