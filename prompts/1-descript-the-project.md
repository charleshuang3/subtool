This is a command line tool to help user edit srt file:

1. Cleanup descriptive subtitles.
2. Analyze if useless subtitle, by calcuating repeat.
3. Cleanup useless subtitle.
4. Translate subtitle via Deepl.

help me write this tool in golang.

the usages of this tool is:

```
subtool remove-descriptive-subtitle -i <input_file>
subtool analyze-repeat -i <input_file>
subtool remove-useless-subtitle -i <input_file> --rm <repeat_file>
subtool translate -i <input_file> --deepl-key <deepl_key_file>
```

please impl functions in lib/sub, and each function should have their own file.
