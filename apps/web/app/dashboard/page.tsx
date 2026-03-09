import { AuthGuard } from "@/components/auth/auth-guard";
import { MainLayout } from "@/components/layouts/main-layout";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { ROUTES } from "@/lib/constants/routes";

export default function DashboardPage() {
  return (
    <AuthGuard>
      <MainLayout>
        <div className="grid gap-6">
          <div className="flex items-center justify-between">
            <div>
              <h1 className="text-2xl font-semibold">Dashboard</h1>
              <p className="text-sm text-fg/70">
                Your home base for work in progress.
              </p>
            </div>
            <Button asChild href={ROUTES.settings} variant="secondary">
              Theme settings
            </Button>
          </div>
          <div className="grid gap-4 md:grid-cols-2">
            {[
              "Tasks overview",
              "Active projects",
              "Upcoming deadlines",
              "Recent activity",
            ].map((title) => (
              <Card key={title}>
                <h3 className="text-lg font-medium">{title}</h3>
                <p className="mt-2 text-sm text-fg/70">
                  Connect your data when you are ready.
                </p>
              </Card>
            ))}
          </div>
        </div>
      </MainLayout>
    </AuthGuard>
  );
}
