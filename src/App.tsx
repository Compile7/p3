import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./App.scss";
import Auth from "./Components/Auth";
import Review from "./Components/Review";
import Team from "./Components/Team";
function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route index path="/" element={<Auth />} />
          <Route path="/add-review" element={<Review />} />
          <Route path="/team" element={<Team />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
