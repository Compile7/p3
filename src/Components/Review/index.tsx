// import "./review.scss";
import Rating from "react-rating";
import Header from "../Header";
import Footer from "../Footer";

const Review = () => {
  const SVGIcon = (props: any) => (
    <svg className={props.className} pointerEvents="none">
      <use xlinkHref={props.href} />
    </svg>
  );
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
              <textarea name="review" id="review"></textarea>
            </div>
            <div className="form-control">
              <label htmlFor="review">Add Rating</label>
            </div>
            <div className="form-control">
              <button type="submit">Submit Review</button>

              <Rating
                emptySymbol={
                  <SVGIcon href="#icon-star-empty" className="icon" />
                }
                fullSymbol={<SVGIcon href="#icon-star-full" className="icon" />}
              />
            </div>
          </form>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default Review;
