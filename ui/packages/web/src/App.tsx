import { Footer, Navigation } from "./components/Layout";
import { AppRoutes } from "./routes";

const App = () => (
  <div className="bg-background flex min-h-screen flex-col">
    <Navigation />
    <div className="flex h-full w-full flex-grow items-center py-4">
      <AppRoutes />
    </div>
    <Footer />
  </div>
);

export default App;
