# golang-metrics-listener

This program in intened to be a univeral reciever of metrics from applications.  The applications are expected to publish their metrics in the following JSON format:

<pre>
<code>
{
    "bu":"xxx (Used as part of automatic routing key)",
    "env":"prod,dev,qa,stage...(Used as part of automatic routing key",
    "hostEnv":"Azure, AWS, DC... (Used as part of automatic routing key, listener will decorate if not given",
    "app":"listener",
    "version":"- optional",
    "route":"p.mos.someextraroute - This is optional",
    "token":"ABCDEFGHIJK (Optional, can be used for security or decoration..)",
    "points": [
        {
        "name":"cpu.idle",
        "value": 61,
        "timestamp":1525462241,
        "tags":[]
        }
    ]
}
</code>
</pre>

Default Routeing Key is going to be bu.env.hostEnv.app but if there is a value in route, this overrides the previous meintioned pairs.

Second part of this project will be to create listeners that extract this info into different tools.
