# My Personal Knowledge Management Tools

This repo contains a set of tools to gather information from various sources to enable task tracking, dashboards, and more.

## Goals

The primary goal for these tools is to enable focus by drawing attention to the most important work items and reducing the effort to do so.

They will do this by...

    * **Gathering data from disparate systems**. Issues can be tracked in Jira and reviews largely come from GitHub. Collect them in one place.
    * **Automatically executing**. Necessary data should be updated automatically and regularly. Summaries of activity that can be inferred should be.
    * **Keeping everything machine readable**. Where it make sense, ensure that data sources are structured so that programs can act on them.
    * **Operating locally**. The canonical sources of information don't exist on my system, but processing should happen locally whenever possible.

## Initial Steps

    * Collecting actionable Jira and GitHub issues/review requests on a daily basis to provide a dashboard each day.
    * Automating daily task tracking as much as possible
        - Observe git repos for branches with names matching assigned Jira or GH issues and provide a report of projects worked
        - Observe GitHub activity for reviews/comments and provide a report of reviews done
    * Propose next work items based on priority and effort required


## Gathering Data

`pkm` is designed to provide an execution environment for small, dedicated scripts. The following environment variables are presented
to scripts:

    * `PKM_DIR`: Defaults to `$HOME/.pkm`
    * `PKM_DATA_DIR`: Where results should be stored. Defaults to `$PKM_DIR/data`.
    * `PKM_SCRIPT_DIR`: Where scripts that retrieve data should be stores. Default sto `$PKM_DIR/scripts`
    * `PKM_GIT_ROOT`: Top-level directory that contains git clones for scripts to search. Defaults to `$HOME/projects`

You can also see the available environment variables by running `pkm env`.
