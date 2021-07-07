package messages

var (
	ERROR_INVALID_CLI_VERSION     = "Astronomer CLI version is not valid"
	ERROR_GITHUB_JSON_MARSHALLING = "Failed to JSON decode Github response from %s"
	ERROR_INVALID_AIRFLOW_VERSION = "Unsupported Airflow Version specified. Please choose from: %s \n"
	ERROR_NEW_MAJOR_VERSION       = "There is an update for Astro CLI. You're using version %s, but %s is the server version.\nPlease upgrade to the matching version before continuing. See https://www.astronomer.io/docs/cli-quickstart for more information.\nTo skip this check use the --skip-version-check flag.\n"

	CLI_CMD_DEPRECATE         = "Deprecated in favor of %s\n"
	CLI_CURR_VERSION          = "Astro CLI Version: %s"
	CLI_CURR_COMMIT           = "Git Commit: %s"
	CLI_CURR_VERSION_DATE     = CLI_CURR_VERSION + " (%s)"
	CLI_LATEST_VERSION        = "Astro CLI Latest: %s"
	CLI_LATEST_VERSION_DATE   = CLI_LATEST_VERSION + " (%s)"
	CLI_INSTALL_CMD           = "\t$ curl -sL https://install.astronomer.io | sudo bash \nOR for homebrew users:\n\t$ brew install astronomer/tap/astro"
	CLI_RUNNING_LATEST        = "You are running the latest version."
	CLI_CHOOSE_WORKSPACE      = "Please choose a workspace:"
	CLI_SET_WORKSPACE_EXAMPLE = "\nNo default workspace detected, you can list workspaces with \n\tastro workspace list\nand set your default workspace with \n\tastro workspace switch [WORKSPACEID]\n\n"
	CLI_UPGRADE_PROMPT        = "A newer version of the Astronomer CLI is available.\nTo upgrade to latest, run:"
	CLI_UNTAGGED_PROMPT       = "Your current Astronomer CLI is not tagged.\nThis is likely the result of building from source. You can install the latest tagged release with the following command"
	CLI_DEPLOYMENT_HARD_DELETE_PROMPT = "\nThis operation is irreversible and permanent. Are you sure?"

	CONFIG_CREATE_DIR_ERROR        = "Error creating config directory\n"
	CONFIG_CREATE_HOME_ERROR       = "Error creating default config in home dir: %s"
	CONFIG_CREATE_FILE_ERROR       = "Error creating config file\n"
	CONFIG_PATH_KEY_MISSING_ERROR  = "Must specify config key\n"
	CONFIG_PATH_KEY_INVALID_ERROR  = "Config does not exist, check your config key\n"
	CONFIG_PROJECT_NAME_ERROR      = "Project name is invalid\n"
	CONFIG_INIT_PROJECT_CONFIG     = "Initialized empty astronomer project in %s"
	CONFIG_INVALID_SET_ARGS        = "Must specify exactly two arguments (key value) when setting a config\n"
	CONFIG_READ_ERROR              = "Error reading config in home dir: %s\n"
	CONFIG_REINIT_PROJECT_CONFIG   = "Reinitialized existing astronomer project in %s\n"
	CONFIG_SAVE_ERROR              = "Error saving config\n"
	CONFIG_SET_DEFAULT_WORKSPACE   = "Default \"%s\" (%s) workspace found, setting default workspace.\n"
	CONFIG_SET_SUCCESS             = "Setting %s to %s successfully\n"
	CONFIG_USE_OUTSIDE_PROJECT_DIR = "You are attempting to %s a project config outside of a project directory\n To %s a global config try\n%s\n"

	COMPOSE_CREATE_ERROR         = "Error creating docker-compose project"
	COMPOSE_IMAGE_BUILDING_PROMT = "Building image..."
	COMPOSE_STATUS_CHECK_ERROR   = "Error checking docker-compose status"
	COMPOSE_STOP_ERROR           = "Error stopping and removing containers"
	COMPOSE_PAUSE_ERROR          = "Error pausing project containers"
	COMPOSE_RECREATE_ERROR       = "Error building, (re)creating or starting project containers"
	COMPOSE_PUSHING_IMAGE_PROMPT = "Pushing image to Astronomer registry"
	COMPOSE_LINK_WEBSERVER       = "Airflow Webserver: http://localhost:%s"
	COMPOSE_LINK_POSTGRES        = "Postgres Database: localhost:%s/postgres"
	COMPOSE_USER_PASSWORD        = "The default credentials are admin:admin"

	ENV_PATH      = "Error looking for \"%s\""
	ENV_FOUND     = "Env file \"%s\" found. Loading...\n"
	ENV_NOT_FOUND = "Env file \"%s\" not found. Skipping...\n"

	HOUSTON_BASIC_AUTH_DISABLED      = "Basic authentication is disabled, conact administrator or defer back to oAuth"
	HOUSTON_CURRENT_VERSION          = "Astro Server Version: %s"
	HOUSTON_DEPLOYMENT_HEADER        = "Authenticated to %s \n\n"
	HOUSTON_DEPLOYING_PROMPT         = "Deploying: %s\n"
	HOUSTON_NO_DEPLOYMENTS_ERROR     = "No airflow deployments found"
	HOUSTON_DEPLOYMENT_NAME_ERROR    = "Please specify a valid deployment name"
	HOUSTON_SELECT_DEPLOYMENT_PROMPT = "Select which airflow deployment you want to deploy to:"
	HOUSTON_OAUTH_REDIRECT           = "Please visit the following URL, authenticate and paste token in next prompt\n"
	HOUSTON_INVALID_DEPLOYMENT_KEY   = "Invalid deployment selection\n"
	// TODO: @adam2k remove this message once the Houston API work is completed that will surface a similar error message
	HoustonInvalidDeploymentUsers = "No users were found for this deployment.  Check the deploymentId and try again.\n"

	INPUT_PASSWORD    = "Password: "
	INPUT_USERNAME    = "Username (leave blank for oAuth): "
	INPUT_OAUTH_TOKEN = "oAuth Token: "

	REGISTRY_AUTH_SUCCESS        = "Successfully authenticated to %s\n"
	RegistryAuthFail             = "\nFailed to authenticate to the registry. Do you have Docker running?\nYou will not be able to push new images to your Airflow Deployment unless Docker is running.\nIf Docker is running and you are seeing this message, the registry is down or cannot be reached.\n"
	REGISTRY_UNCOMMITTED_CHANGES = "Project directory has uncommmited changes, use `astro deploy [releaseName] -f` to force deploy."

	SETTINGS_PATH = "Error looking for settings.yaml"

	NA                          = "N/A"
	VALID_DOCKERFILE_BASE_IMAGE = "quay.io/astronomer/ap-airflow"
	WARNING_DOWNGRADE_VERSION   = "Your Astro CLI Version (%s) is ahead of the server version (%s).\nConsider downgrading your Astro CLI to match. See https://www.astronomer.io/docs/cli-quickstart for more information.\n"
	WARNING_INVALID_IMAGE_NAME  = "WARNING! The image in your Dockerfile is pulling from '%s', which is not supported. We strongly recommend that you use Astronomer Certified images that pull from 'astronomerinc/ap-airflow' or 'quay.io/astronomer/ap-airflow'. If you're running a custom image, you can override this. Are you sure you want to continue?\n"
	WARNING_INVALID_IMAGE_TAG   = "WARNING! You are about to push an image using the '%s' tag. This is not recommended.\nPlease use one of the following tags: %s.\nAre you sure you want to continue?"
	WARNING_NEW_MINOR_VERSION   = "A new minor version of Astro CLI is available. Your version is %s and %s is the latest.\nSee https://www.astronomer.io/docs/cli-quickstart for more information.\n"
)
