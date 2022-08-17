# TODO

 - [x] Make a sub-directory by day
 - [x] Make cache list work
 - [x] Should check file for write protected
 - [x] Should check all files in a directory for write protected
 - [x] List output should have color
 - [x] Should say if list output file is a directory
 - [x] Should show how many files are in trash list, like: .../trash/21/foo [21 files]
 - [x] When listing cache, fix issue when theres less then 10 files there
 - [x] Fix issue when removing a broken link
 - [x] Add tests
 - [ ] Add more tests
 - [ ] Add option to disable color
 - [x] Recover option should be able to accept more then one number
 - [x] Should also autoremove files under 100 bytes
 - [x] Make sure only 1 file per trash index path
 - [x] Better format of bytes (autoclean)
 - [ ] Use better sorting func defining
 - [ ] Autocleaning should have an option to remove binary (non ascii) files
 - [x] Caching should make sub-directories every day (eg. ~/.cache/srm/trash/2022/04/20/[0,1,2,3,4,5,6,7,8...])
 - [ ] Autocleaning should remove month-old files
 - [ ] Eventally, autocleaning should run if needed during running if total space is grader than an amount
 - [ ] Should keep a track file of the total space used in a ini format so users can also see it (this way it will
       be easy to know when to run the autocleaning, see todo item above)
 - [ ] Should have a `SRM_TRASH` env var or something to specify the trash path
 - [x] Should support the `--` for end of flags and start of arguments (already handled by cli library)

