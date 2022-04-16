# configMaps

a configMap is an api object used to store non-confidential data in a key-value pairs. Pods can consume configMaps as environment variables, command-line arguments, or as configuration files in a volume.

a configMap allows you to decouple evironment-specific configuration from your container-images, so that your application are easily portable.

> configMap does not provide secrecy or encyption. if the data you want to store are confidential, use a secret rather than a configMap, or use additional (third party) tools to keep your data private.
