import { useState } from "react";
import { useStateContext } from "../Contexts/contextProvider";
import { useNavigate } from "react-router-dom";

const useFetch = (url:string) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");
  const navigate = useNavigate();
  const { setProfile} = useStateContext();
  const handleGoogle = async (response : any) => {
    setLoading(true);
    console.log(response)
    fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization" : `Bearer ${response.credential}`
      },
    })
      .then((res) => {
        setLoading(false);
        
        return res.json();
      })
      .then((data) => {
        console.log(data)
        if (data.Email) {
          localStorage.setItem("P3AccessToken", response.credential);
          setProfile(data)
          
          data.Role === 'Admin'?
          navigate('/onboarding') : 
          navigate("/view-review")
        }
        // throw new Error(data?.message || data);
      })
      .catch((error) => {
        setError(error?.message);
      });
  };
  return { loading, error, handleGoogle };
};

export default useFetch;