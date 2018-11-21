import boto3
import constants
import sys

client = boto3.client('lambda')

func = None

for currentFunc in client.list_functions():
    if currentFunc.FunctionName == constants.FUNCTION_NAME:
        func = currentFunc
        break

if func is None:
    sys.exit('Error!')

print 'FunctionName:%0'.format(func.FunctionName)
print 'FunctionName:%0'.format(func.FunctionArn)

# test for function existence

# if function doesn't exist
# - create skeleton
# - upload code

# - update code


# response = client.create_function(
#     FunctionName='TeslaIotKey',
#     Runtime='go1.x',
#     Role='string',
#     Handler='Handle',
#     Code={
#         'ZipFile': b'bytes',
#         'S3Bucket': 'string',
#         'S3Key': 'string',
#         'S3ObjectVersion': 'string'
#     },
#     Description='Tesla IOT Key',
#     Timeout=90,
#     MemorySize=128,
#     Publish=True,
#     DeadLetterConfig={
#         'TargetArn': 'string'
#     },
#     Environment={
#         'Variables': {
#             'string': 'string'
#         }
#     },
#     KMSKeyArn='string',
#     Tags={
#         'author': 'stephbu@gmail.com'
#     }
# )
