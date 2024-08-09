from typing import Dict

def str2bool(string: str) -> bool:
  return string.lower() in ('yes', 'true', '1', 'y', 't')

def error_dict(message: str) -> Dict[str, str]:
  return {
    'error': message,
  }
