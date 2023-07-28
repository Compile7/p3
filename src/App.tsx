import { BrowserRouter, Routes, Route, Navigate } from "react-router-dom";
import "./App.scss";
import Auth from "./Components/Auth";
import Review from "./Components/Review";
import ViewReview from "./Components/ViewReview";
import Team from "./Components/Team";
// import { ProtectedRoute } from "./Components/ProtectedRoute";
import { isLoggedIn } from "./utils";
import useFetch from "./Hooks/useFetch";
import Onboarding from "./Components/Onboarding";
import { useEffect } from "react";
import { ProtectedRoute } from "./Components/ProtectedRoute";
import Invitations from "./Components/Invitations";
// import { useStateContext } from "./Contexts/contextProvider";
const App=()=> {
  const { handleGoogle } = useFetch(`${import.meta.env.VITE_APP_BACKEND_ENDPOINT}/login`);
    
    useEffect(() => {
      console.log('reach')
      if (isLoggedIn()) {
        console.log("Working")
        
        handleGoogle({ credential: localStorage.getItem("P3AccessToken") });
      }
    
      return () => {
       
      }
    },[])
    console.log('rea2ch')
  
  return (
    <>
      <BrowserRouter>
        <Routes>
          <Route index path="/" element={<Auth />} />
          <Route
            path="/:name/add-review"
            element={
              <ProtectedRoute roleProtected>
                <Review />
               </ProtectedRoute>
            }
          />
          <Route
            path="/invitations"
            element={
              <ProtectedRoute roleProtected>
                <Invitations />
              </ProtectedRoute>
            }
          />
          <Route
            path="/team"
            element={
              <ProtectedRoute roleProtected>
                <Team />
              </ProtectedRoute>
            }
          />
          <Route
            path="/view-review"
            element={
              <ProtectedRoute>
                <ViewReview />
              </ProtectedRoute>
            }
          />
          <Route
            path="/onboarding"
            element={
              <ProtectedRoute roleProtected>
                <Onboarding />
              </ProtectedRoute>
            }
          />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </BrowserRouter>
    </>
  );
}

export default App;