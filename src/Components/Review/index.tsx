import Header from "../Header";
import Footer from "../Footer";
import { useParams } from "react-router";
import { useState } from "react";

const Review = () => {

  const params = useParams()
  console.log(params.name)

  const [loading, setLoading] = useState(false)

  const addReview = ()=>{

    setLoading(true)

    //api call for review

    console.log("API call for add review")
    setLoading(false)

  }
  return (
    <>
      <Header />
      <main>
        <div>
          <div className="header">
            <h1>{`Add Review for ${params.name}`}</h1>
            <p>Fill the form to submit review</p>
          </div>
          <form action="">
            <div className="form-control">
              <label htmlFor="review">Enter Review</label>
              <textarea name="review" id="review" required></textarea>
            </div>
            <div className="form-control">
              <label htmlFor="review">Add Rating</label>
              <input type="number" max={10} min={1} />
            </div>
            <div className="form-control">
              <button type="submit" className="btn btn-primary" onClick={addReview}>
                Submit Review
              </button>
            </div>
          </form>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default Review;
