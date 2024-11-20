import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { LoginPage, RegisterPage } from "./features/auth/views";

const router = createBrowserRouter([
  {
    path: "/",
    element: <LoginPage />,
  },
  {
    path: "/register",
    element: <RegisterPage />,
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
