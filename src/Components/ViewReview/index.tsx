import Header from "../Header";
import Footer from "../Footer";

const ViewReview = () => {
  return (
    <>
      <Header />
      <main>
        <div>
          <div className="header has-selement">
            <div>
              <h1>Hitesh Kumawat</h1>
              <p>Product Designer</p>
            </div>
            <div>
              <select className="form-select primary" id="month" name="month">
                <option value="">Select month</option>
                <option value="01">January</option>
                <option value="02">February</option>
                <option value="03">March</option>
                <option value="04">April</option>
                <option value="05">May</option>
                <option value="06">June</option>
                <option value="07">July</option>
                <option value="08">August</option>
                <option value="09">September</option>
                <option value="10">October</option>
                <option value="11">November</option>
                <option value="12">December</option>
              </select>
            </div>
          </div>
          <div className="review-container">
            <div className="month">
              <span className="badge badge-primary">July month</span>
            </div>
            <div className="reviewer-detail">
              <div className="reviewername">Abhimanyu Singh</div>
              <span className="badge badge-wraning">Reviewer</span>
            </div>
            <div className="reviews">
              <p>
                Lorem Ipsum is simply dummy text of the printing and typesetting
                industry.
              </p>
              <ul>
                <li>
                  Lorem Ipsum is simply dummy text of the printing and
                  typesetting industry.
                </li>
                <li>
                  Lorem Ipsum is simply dummy text of the printing and
                  typesetting industry.
                </li>
                <li>
                  Lorem Ipsum is simply dummy text of the printing and
                  typesetting industry.
                </li>
                <li>
                  Lorem Ipsum is simply dummy text of the printing and
                  typesetting industry.
                </li>
              </ul>
            </div>
          </div>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default ViewReview;
