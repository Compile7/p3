import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./App.scss";
import Auth from "./Components/Auth";
import Review from "./Components/Review";
import ViewReview from "./Components/ViewReview";
import Team from "./Components/Team";
import { ProtectedRoute } from "./Components/ProtectedRoute";
import { isLoggedIn } from "./utils";
import useFetch from "./Hooks/useFetch";
// import { useStateContext } from "./Contexts/contextProvider";
function App() {

   if (isLoggedIn()){
    const { handleGoogle } = useFetch(
      "http://localhost:5152/login"
    );
    handleGoogle({credential: localStorage.getItem("P3AccessToken")});
  
   }
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route index path="/" element={<Auth />} />
          <Route path="/add-review" element={<ProtectedRoute ><Review /></ProtectedRoute>} />
          <Route path="/team" element={<ProtectedRoute ><Team /></ProtectedRoute>} />
          <Route path="/view-review" element={<ProtectedRoute ><ViewReview /></ProtectedRoute>} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
