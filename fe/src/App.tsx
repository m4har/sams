import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { LoginPage } from "./features/auth/views";

const router = createBrowserRouter([
  {
    path: "/",
    element: <LoginPage />,
  },
]);
const App = () => {
  return (
    <>
      <RouterProvider router={router} />
    </>
  );
};

export default App;
