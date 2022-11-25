import Header from "./header"
import Login from "./login"

export default function Home() {
    return (
      <>
        <Header/>
        <div className="flex flex-auto justify-center items-center">
          <Login />
        </div>
      </>
    );
  }