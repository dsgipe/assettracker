## Asset Tracker

### Performance Notes
1. The first commit shows a pretty basic approach using maps and the encoding/json scanner.
   1. Performance is not great, about 30% slower than when nothing is done aside
    from scanning the input.
2. The second commit refactors the map to use an array with the market as index. This improves performance by 15%
3. The next bottleneck came from unmarshalling.
   1. In order to further improve performance, I investigated doing all unmarshalling at the end of the run. However, 
   this worsened performance significantly
   2. I then evaluated multiple unmarshalers and their performance.
      1. github.com/json-iterator/go
         1. showed some improvement, but still seemed worth looking at others
      2. github.com/mailru/easyjson/
         1. showed significant improvement and enough to satisfy my goal of determining if there is a better way.
         2. It seems that the code generation is worth the performance improvement over reflection. 
         3. Whether this is best option seems to be out of scope.
4. I also looked at adding goroutines. However, the performance overhead significantly outweighed any benefit
      
## A note on maps vs arrays
I implemented both maps and arrays
1. Maps directly were about 15% slower than using arrays
2. Pre-allocating arrays was faster than my eventual decision on using a map to map between indices
3. My final decision to implement a map around the index was to provide a more robust solution in case the markets were 
not in 0-max number format, and only increased the time by a marginal amount. Around 1 percent. Also there is the benefit 
to memory allocation and the eventual garbage collector if not all indices are filled.
4. I think maps are the easier to read solution, however for performance gains I am happy with using arrays and mapping
the index.

## Easyjson
1. `run easyjson manageincoming/models.go` to generate json unmarshaler code

