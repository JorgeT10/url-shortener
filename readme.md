# **Exercise: URL Shortener**

 You are tasked with implementing a basic URL shortener in Go.
 The shortener should be able to shorten a given URL and redirect to the original URL when the shortened version is accessed.
 

#### Here are the basic requirements:

  - Create a function ShortenURL(originalURL string) string that takes an
   original URL and returns a shortened version.
   
 -  Create a function Redirect(shortURL string) (string, error) that
   takes a shortened URL and returns the original URL. If the shortened
   URL is not found, return an error.   
   - Use an in-memory data structure to store the mapping between
   shortened and original URLs.   
   - Ensure that the shortened URLs are unique.   
   - Implement basic error handling.
   - Collision Resistance: Add a collision detection and resolution mechanism to handle cases when the generated short URL already exists.
   - Hit Count Tracking: Track the number of times a short URL is accessed, providing basic analytics.
   - Thread Safety: Make the URLShortener struct thread-safe for concurrent access.

#### Tips for the interview

- Consider writing test scenarios and validating them with the interviewer.
- The short URL should not have a valid domain, just a fake domain will be fine.
