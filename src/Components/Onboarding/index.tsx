import Header from "../Header";
import Footer from "../Footer";
import onboarding from "../../assets/onboarding.png";
const Onboarding = () => {
  return (
    <>
      <Header />
      <main>
        <div>
          {/* Step 1 */}
          <div className="onboarding">
            <div className="description">
              <div className="header">
                <h1>Create a new Organization</h1>
                <p>
                  Lorem Ipsum is simply dummy text of the printing and
                  typesetting industry. Lorem Ipsum has been the industry's
                  standard dummy text ever since the 1500s
                </p>
              </div>
              <div>
                <a href="#" className="btn btn-primary">
                  Create organization
                </a>
              </div>
            </div>
            <div className="image">
              <img src={onboarding} alt="onboarding" />
            </div>
          </div>
          {/* Step 2 */}
          <div className="onboarding">
            <div className="description">
              <div className="steps">
                Step <span>1</span> of 2
              </div>
              <div className="header">
                <h1>What's the name of your organization</h1>
                <p>
                  Lorem Ipsum is simply dummy text of the printing and
                  typesetting industry.
                </p>
              </div>
              <div>
                <form action="">
                  <div className="form-control">
                    <label htmlFor="organization_name">Organization name</label>
                    <input type="text" id="organization_name" />
                  </div>
                  <div className="form-control">
                    <button type="submit" className="btn btn-primary">
                      Next
                    </button>
                  </div>
                </form>
              </div>
            </div>
            <div className="image">
              <img src={onboarding} alt="onboarding" />
            </div>
          </div>

          {/* Step 3 */}
          <div className="onboarding">
            <div className="description">
              <div className="steps">
                Step <span>2</span> of 2
              </div>
              <div className="header">
                <h1>Invite your team</h1>
                <p>
                  Lorem Ipsum is simply dummy text of the printing and
                  typesetting industry.
                </p>
              </div>
              <div>
                <div className="import-data">
                  <div className="help-text">
                    <svg
                      width="48"
                      height="48"
                      viewBox="0 0 48 48"
                      fill="none"
                      xmlns="http://www.w3.org/2000/svg"
                    >
                      <g clip-path="url(#clip0_32_166)">
                        <path
                          d="M40 12V36C40 37.0609 39.5786 38.0783 38.8284 38.8284C38.0783 39.5786 37.0609 40 36 40H16C14.9391 40 13.9217 39.5786 13.1716 38.8284C12.4214 38.0783 12 37.0609 12 36V12C12 10.9391 12.4214 9.92172 13.1716 9.17157C13.9217 8.42143 14.9391 8 16 8H36C37.0609 8 38.0783 8.42143 38.8284 9.17157C39.5786 9.92172 40 10.9391 40 12Z"
                          stroke="#64748B"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                        <path
                          d="M20 32H32"
                          stroke="#64748B"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                        <path
                          d="M22 22C22 23.0609 22.4214 24.0783 23.1716 24.8284C23.9217 25.5786 24.9391 26 26 26C27.0609 26 28.0783 25.5786 28.8284 24.8284C29.5786 24.0783 30 23.0609 30 22C30 20.9391 29.5786 19.9217 28.8284 19.1716C28.0783 18.4214 27.0609 18 26 18C24.9391 18 23.9217 18.4214 23.1716 19.1716C22.4214 19.9217 22 20.9391 22 22Z"
                          stroke="#64748B"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                        <path
                          d="M8 16H14"
                          stroke="#64748B"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                        <path
                          d="M8 24H14"
                          stroke="#64748B"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                        <path
                          d="M8 32H14"
                          stroke="#64748B"
                          stroke-width="2"
                          stroke-linecap="round"
                          stroke-linejoin="round"
                        />
                      </g>
                      <defs>
                        <clipPath id="clip0_32_166">
                          <rect width="48" height="48" fill="white" />
                        </clipPath>
                      </defs>
                    </svg>
                    <span>Quickly invite from your Google contact list</span>
                  </div>
                  <div className="google-btn">
                    <button>
                      <div className="g-sign-in-button">
                        <div className="logo-wrapper">
                          <svg
                            width="20"
                            height="20"
                            viewBox="0 0 20 20"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                          >
                            <path
                              d="M10.0002 3.95833C11.4752 3.95833 12.7961 4.46667 13.8377 5.45833L16.6919 2.60417C14.9586 0.991667 12.6961 0 10.0002 0C6.09189 0 2.71273 2.24167 1.06689 5.50833L4.39189 8.0875C5.17939 5.71667 7.39189 3.95833 10.0002 3.95833Z"
                              fill="#EA4335"
                            />
                            <path
                              d="M19.575 10.2291C19.575 9.57498 19.5125 8.94165 19.4167 8.33331H10V12.0916H15.3917C15.15 13.325 14.45 14.375 13.4 15.0833L16.6208 17.5833C18.5 15.8416 19.575 13.2666 19.575 10.2291Z"
                              fill="#4285F4"
                            />
                            <path
                              d="M4.3875 11.9125C4.1875 11.3084 4.07083 10.6667 4.07083 10C4.07083 9.33336 4.18333 8.69169 4.3875 8.08753L1.0625 5.50836C0.383333 6.85836 0 8.38336 0 10C0 11.6167 0.383333 13.1417 1.06667 14.4917L4.3875 11.9125Z"
                              fill="#FBBC05"
                            />
                            <path
                              d="M10 20C12.7 20 14.9708 19.1125 16.6208 17.5792L13.4 15.0792C12.5042 15.6833 11.35 16.0375 10 16.0375C7.39167 16.0375 5.17917 14.2792 4.3875 11.9083L1.0625 14.4875C2.7125 17.7583 6.09167 20 10 20Z"
                              fill="#34A853"
                            />
                          </svg>
                        </div>
                        <span className="text-container">
                          <span>Import contacts</span>
                        </span>
                      </div>
                    </button>
                  </div>
                </div>
                <form action="">
                  <div className="form-control">
                    <label htmlFor="review">Add team members</label>
                    <textarea
                      name="review"
                      id="review"
                      required
                      placeholder="Ex. hitesh@loginradius.com, ankit@loginradius.com"
                    ></textarea>
                  </div>
                  <div className="form-control">
                    <div className="btn-group">
                      <button type="submit" className="btn btn-primary">
                        Add members
                      </button>
                      <button type="submit" className="btn btn-ghost">
                        Skip this step
                      </button>
                    </div>
                  </div>
                </form>
              </div>
            </div>
            <div className="image">
              <img src={onboarding} alt="onboarding" />
            </div>
          </div>
        </div>
      </main>
      <Footer />
    </>
  );
};

export default Onboarding;
