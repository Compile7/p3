import { useState } from "react";
import { useStateContext } from "../Contexts/contextProvider";

const useFetch = (url:string) => {
  const [loading, setLoading] = useState(false);
  const [error, setError] = useState("");

  const { setProfile} = useStateContext();
  const handleGoogle = async (response : any) => {
    setLoading(true);
    console.log(response)
    fetch(url, {
      method: "GET",
      headers: {
        "Content-Type": "application/json",
        "Authorization" : ` Bearer ${response.credential}`
      },
    })
      .then((res) => {
        setLoading(false);
        
        return res.json();
      })
      .then((data) => {
        if (data?.user) {
          localStorage.setItem("P3AccessToken", response.credential);
          setProfile(data)
          window.location.reload();
        }

        throw new Error(data?.message || data);
      })
      .catch((error) => {
        setError(error?.message);
      });
  };
  return { loading, error, handleGoogle };
};

export default useFetch;