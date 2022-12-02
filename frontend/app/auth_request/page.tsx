import Description from './description';
import RequestForm from './form';

export default function Home() {
    return (
      <div className="flex flex-auto justify-center items-center h-full">
        <div className="h-full grid grid-cols-1 md:grid-cols-2">
          <div className="p-10 lg:px-20 md:self-center">
            <Description />
          </div>
          <div className="p-10 lg:px-20 md:self-center">
            <RequestForm />
          </div>
        </div>
      </div>
    );
  }