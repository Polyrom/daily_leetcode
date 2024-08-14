## Static & Dynamic Arrays, Strings

[AlgoMap Youtube clip](https://www.youtube.com/watch?v=TQMvBTKn2p0)

### Static Arrays

- Access/modify element: O(1)
- Check element's presence: O(n)
- Insert element: O(n)
- Delete element: O(n)

### Dynamic Arrays

- Access/modify element: O(1)
- Check element's presence: O(n)
- Insert element, not at the end: O(n)
- Append: \*O(1) (O(n) if array length surpassed)
- Delete element, not at the end: O(n)
- Popping element at the end: O(1)

### Strings

- Access element: O(1)
- All other operations are O(n), as strings are immutable (at least in Python)
- Due to immutable nature, modifying is impossible
