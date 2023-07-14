import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./App.scss";
import Auth from "./Components/Auth";
import Review from "./Components/Review";
import ViewReview from "./Components/ViewReview";
import Team from "./Components/Team";
import { ProtectedRoute } from "./Components/ProtectedRoute";
function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route index path="/" element={<Auth />} />
          <ProtectedRoute path="/add-review" element={<Review />} />
          <ProtectedRoute path="/team" element={<Team />} />
          <ProtectedRoute path="/view-review" element={<ViewReview />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
