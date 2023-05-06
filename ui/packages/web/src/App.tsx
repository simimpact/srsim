import { Footer, Navigation } from "./components/Layout";
import { AppRoutes } from "./routes";

const App = () => (
  <div className="h-screen flex flex-col">
    <Navigation />
    <div className="w-full flex flex-col items-center flex-grow pb-4">
      <AppRoutes />
    </div>
    <Footer />
  </div>
);

export default App;
