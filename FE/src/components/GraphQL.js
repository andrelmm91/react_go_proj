import { useEffect, useState } from "react";
import { Link } from "react-router-dom";
import Input from './form/Input';

const GraphQL = () => {
    // set up stateful variables
    const [movies, setMovies] = useState([]);
    const [searchTerm, setSearchTerm] = useState("");
    const [fullList, setFullList] = useState([]);

    // perform a search
    const performSearch = () => {
        const payload = `
        {
            search(titleContains: "${searchTerm}") {
                id
                title
                runtime
                release_date
                mpaa_rating
            }
        }`;

        const headers = new Headers();
        headers.append("Content-Type", "application/graphql");

        const requestOptions = {
            method: "POST",
            body: payload,
            headers: headers,
        }

        fetch(`/graph`, requestOptions)
            .then((response) => response.json())
            .then((response) => {
                let theList = Object.values(response.data.search);
                setMovies(theList);
            })
            .catch(err => {console.log(err)})
    }

    const handleChange = (event) => {
        event.preventDefault();

        let value = event.target.value;
        setSearchTerm(value);

        if (value.length > 2) {
            performSearch();
        } else {
            setMovies(fullList);
        }
    }

    // useEffect
    useEffect(() => {
        const payload = `
        {
            list {
                id
                title
                runtime
                release_date
                mpaa_rating
            }
        }`;

        const headers = new Headers();
        headers.append("Content-Type", "application/grapql");

        const requestOptions = {
            method: "POST",
            headers: headers,
            body: payload,
        }

        fetch(`/graph`, requestOptions)
            .then((response) => response.json())
            .then((response) => {
                let theList = Object.values(response.data.list);
                setMovies(theList);
                setFullList(theList);
            })
            .catch(err => {console.log(err)})
    }, [])

    return(
        <div>
            <h2>GraphQL</h2>
            <hr />

            <form onSubmit={handleChange}>
                <Input
                    title={"Search"}
                    type={"search"}
                    name={"search"}
                    className={"form-control"}
                    value={searchTerm}
                    onChange={handleChange}/>
            </form>

            {movies ? (
                <table className="table table-striped table-hover">
                    <thead>
                        <tr>
                            <th>Movie</th>
                            <th>Release Date</th>
                            <th>Rating</th>
                        </tr>
                    </thead>
                    <tbody>
                        {movies.map((m) => (
                            <tr key={m.id}>
                                <td>
                                    <Link to={`/movies/${m.id}`}>
                                        {m.title}
                                    </Link>
                                </td>
                                <td>{new Date(m.release_date).toLocaleDateString()}</td>
                                <td>{m.mpaa_rating}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            ) : (
                <p>No movies (yet)!</p>
            )}
        </div>
    )
}

export default GraphQL;