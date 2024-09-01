import os

# Set the directory path
directory = '/home/dn/cheetah/prod/dnos_monolith/yangs/'

# List all files in the directory
files = os.listdir(directory)

# Filter files that end with .yang and do not contain "private" in their name
filtered_files = [file for file in files if file.endswith('.yang') and 'private' not in file.lower() and 'openconfig' not in file.lower()
                  and 'ietf' not in file.lower() and 'iana' not in file.lower() and 'rpc' not in file.lower() and 'notification' not in file.lower()]

# Print the filtered files
for file in filtered_files:
    print('/dn/yangs/'+file)
