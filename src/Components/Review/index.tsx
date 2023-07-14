import Header from "../Header";
import Footer from "../Footer";

const Review = () => {
  return (
    <>
      <Header />
      <main>
        <div>
          <div className="header">
            <h1>Add Review</h1>
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
              <button type="submit">Submit Review</button>
            </div>
          </form>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default Review;
