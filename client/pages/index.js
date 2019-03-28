import React, { useState, useEffect } from "react";
import ReactDOM from "react-dom";
import axios from 'axios';


function SearchResults() {
  const [data, setData] = useState('none');
  const [query, setQuery] = useState('https://www.tripadvisor.fr/Restaurant_Review-g60763-d1236281-Reviews-Club_A_Steakhouse-New_York_City_New_York.html');
  const [search, setSearch] = useState('');
  const [isLoading, setIsLoading] = useState(false);
  const [isError, setIsError] = useState(false);

  useEffect(() => {
    let ignore = false;
    async function fetchData() {
      setIsError(false);
      setIsLoading(true);
      try {
        const result = await axios.get('http://localhost:8083/getTA?url_name='+query,
          JSON.stringify({ data: 'data' }));
        if (!ignore) setData(result.data);
        setIsLoading(false);
      } catch (error) {
        setIsError(true);
      }

    }

    fetchData();
    return () => { ignore = true; }
  }, [search]);

  return (
    <div>
    <div>
      <input name="url_name" value={query} onChange={e => setQuery(e.target.value)} />
</div><div>
      <input type="button" value="Quick search !" onClick={() => setSearch(query)} />
</div>
      {isError && <div>Error : Couldn't find informations on this page.</div>}

      <div>
      <ul>
        <li>
            Number of comments :
            {isLoading ? (
              <span>Loading ...</span>
            ) : (<span>{data.Nb_comments}</span>)}
          </li><li>
            Rating : {isLoading ? (
                <span>Loading ...</span>
              ) : (<span>{data.Rating}</span>)}
          </li>
        </ul>
      </div>
    </div>
  );
}

// const rootElement = document.getElementById("root");
// ReactDOM.render(<SearchResults />, rootElement);
export default SearchResults;
