
# Bolt Inspector 2
Revised version of [BoltInspector](https://github.com/89yoyos/BoltInspector) for [BBolt](https://github.com/etcd-io/bbolt)

## About
When I was working with Bolt last year, I made [BoltInspector](https://github.com/89yoyos/BoltInspector) to help debug database files. However the main branch of [Bolt](https://github.com/boltdb/bolt) is now no longer maintained, and the [BBolt](https://github.com/etcd-io/bbolt) fork is the suggested branch.

Rather than update [BoltInspector](https://github.com/89yoyos/BoltInspector), which might break compatibility with original Bolt files as [BBoltDB](https://github.com/etcd-io/bbolt) develops, I decided to leave [BoltInspector](https://github.com/89yoyos/BoltInspector) up in its current state and create an 2.0 version in a different repository. I also split out the code I used to manipulate the database into a separate repository as well, now known as [Wrench](https://github.com/89yoyos/Wrench)

## Features
*Features in the most current version*
- [x] Open Database files
- [x] Move between buckets
- [x] List everything in the current bucket
- [x] Recursively list everything in the current bucket and nested buckets
- [x] Distinguish between values and nested buckets when listing
- [x] Print values (String, Int, or UInt)
- [x] Insert new values/buckets
- [x] Move/Copy existing values
- [x] Edit existing values
- [x] Delete values/buckets
- [x] Empty bucket (delete all values/buckets inside a bucket, but not the bucket itself)

## Usage
This is a Command-Line program, and its usage is pretty straight-forward. If you need help with any of the commands, the Help command works wonders. Though, I may add that documentation here at some point.

If the help command doesn't answer your questions, the old documentation is available in the [Wiki of the original repository.](https://github.com/89yoyos/BoltInspector/wiki). I haven't changed much, so it *should* still be accurate. This is a port from the original, however. So while I've tested it a fair bit, I'm sure I missed some edge cases. I'll fix issues as I find them. Feel free to post pull requests or issues here.

## Links

This project is covered under the MIT License. Details can be found [here](./LICENSE)
