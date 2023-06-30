import { BrowserRouter, Routes, Route } from "react-router-dom";
import "./App.scss";
import Auth from "./Components/Auth";
import Review from "./Components/Review";
function App() {
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route index path="/" element={<Auth />} />
          <Route path="/add-review" element={<Review />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;
