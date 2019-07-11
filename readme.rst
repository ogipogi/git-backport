*****
git backport
*****

|Website goreportcard.com|

.. |Website goreportcard.com| image:: https://goreportcard.com/badge/github.com/pjhampton/git-backport
   :target: https://goreportcard.com/report/github.com/pjhampton/git-backport
   
.. image:: https://circleci.com/gh/pjhampton/git-backport.svg?style=svg
    :target: https://circleci.com/gh/pjhampton/git-backport

Need to go back to the future?

``git backport`` is a git plugin that will help backport commits in time to other release branches.

Installing
**********************

`Download the binary <https://github.com/pjhampton/git-backport/releases>`_ and add to your PATH: ``/usr/local/bin/``

Once this is done refresh your terminal and test with ``git backport``

Using
**********************

Specify the commit and the branch(es) that you want to back port to - separated with a colon.

``$ git backport commit_hash:branch_1,branch_2,...,branch_n``

Example:

``$ git backport c972b462d82aa69ee6876fae7d865b9b58e26abe:version_1,version_2,version_3``

Contributing
**********************

Feel free to contribute code - please make sure it's adequately tested.

Unlicense-d
**********************

If it feels good - do it, baby.
