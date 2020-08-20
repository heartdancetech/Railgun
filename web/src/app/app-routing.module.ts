import { NgModule } from "@angular/core";
import { Routes, RouterModule } from "@angular/router";
import { WelcomeComponent } from "./pages/welcome/welcome.component";

const routes: Routes = [
  { path: "", pathMatch: "full", redirectTo: "/welcome" },
  { path: "welcome", component: WelcomeComponent },
  {
    path: "config",
    loadChildren: () =>
      import("./pages/config/config.module").then((m) => m.ConfigModule),
  },
];

@NgModule({
  imports: [RouterModule.forRoot(routes)],
  exports: [RouterModule],
})
export class AppRoutingModule {}
