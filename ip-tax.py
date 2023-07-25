import did.cli
from datetime import date
from dateutil.relativedelta import relativedelta as delta

since = (date.today() - delta(months=1)).strftime('%Y-%m-25')
until = date.today().strftime('%Y-%m-25')

args = [
  '--since', since,
  '--until', until,
  '--github.com-issues-commented',
  '--github.com-pull-requests-created',
  '--gitlab.cee-merge-requests-created',
  '--gitlab.mw-merge-requests-created'
]

# print(args)

did.cli.main(args)
