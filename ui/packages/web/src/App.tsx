import { Footer, Navigation } from "./components/Layout";
import { AppRoutes } from "./routes";

const App = () => (
  <div className="min-h-screen flex flex-col bg-background">
    <Navigation />
    <div className="w-full h-full flex items-center flex-grow py-4">
      <AppRoutes />
    </div>
    <Footer />
  </div>
);

export default App;
