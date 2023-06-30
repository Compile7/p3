import Logo from "../../assets/logo.svg";
const Header = () => {
  return (
    <header>
      <div>
        <div className="logo">
          <a href="/">
            <img src={Logo} alt="Logo" />
          </a>
        </div>
        <nav>
          <ul>
            <li>
              <a href="#">Menu01</a>
            </li>
            <li>
              <a href="#">Menu02</a>
            </li>
          </ul>
        </nav>
      </div>
    </header>
  );
};
export default Header;
