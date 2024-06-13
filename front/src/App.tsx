import type { JSXElement } from "solid-js";
import { Router, Route } from "@solidjs/router";
import HomePage from "./pages/home";

export default function App(): JSXElement {
  return (
    <Router>
      <Route path="/" component={HomePage} />
    </Router>
  );
}
