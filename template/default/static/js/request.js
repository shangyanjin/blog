/**
 * Create an axios instance with default configuration
 * - Sets base URL for API requests
 * - Configures default headers including auth token and content type
 */
const token = localStorage.getItem('token');
const tokenAdmin = localStorage.getItem('adminToken');
const axiosInstance = axios.create({
    baseURL: '/api',
    headers: {

        'Authorization': `Bearer ${token}`,
        'token': tokenAdmin,
        'Content-Type': 'application/json'

    }
});

/**
 * Add response interceptor
 * - Handles error responses
 * - Standardizes error format
 */
axiosInstance.interceptors.response.use(
    (response) => {
        const { data } = response;
        
        // Only handle code error cases
        if (data.code !== undefined && data.code !== 0 && data.code !== 200) {
            // Handle unauthorized access
            if (data.code === 401 || data.code === 403) {
                handleUnauthorizedAccess();
            }
            console.error(`API Error: [${data.code}] ${data.message || 'API request failed'}`);
            return response;
        }
        
        return response;
    },
    (error) => {
        // Handle network errors or other axios errors
        const message = error.response?.data?.message || error.message || 'Network error';
        const customError = new Error(message);
        customError.code = error.response?.status || 500;
        throw customError;
    }
);

/**
 * Handle unauthorized access by redirecting to appropriate login page
 * - Redirects to /admin/login for admin paths
 * - Redirects to /login for other paths
 * - Prevents redirect loops by checking current location
 */
function handleUnauthorizedAccess() {
   
  
    // Check if not already on a login page
    if (!window.location.href.includes("/login")) {
        let loginPath = "/login";  // default login path
        
        // Get current path
        const currentPath = window.location.pathname;
        
        // Determine specific login path based on current URL
        if (currentPath.startsWith('/admin')) {
            loginPath = '/admin/login';
        }
        
        // Redirect only if not already on the target login page
        if (!currentPath.includes(loginPath)) {
            console.log('Redirecting to login page:', loginPath);
            window.location.href = loginPath;
        }
    }
}

/**
 * Global request object providing HTTP methods
 * Exposes methods for making API requests with pre-configured axios instance
 */
window.request = {
    /**
     * Make a GET request to the specified URL
     * @param {string} relativeUrl - The endpoint URL relative to base URL
     * @param {Object} params - Query parameters to include in the request
     * @returns {Promise<any>} Response data from the server
     */
    async get(relativeUrl, params = {}) {
        const response = await axiosInstance.get(relativeUrl, { params });
        return response.data;
    },

    /**
     * Make a POST request to the specified URL
     * @param {string} relativeUrl - The endpoint URL relative to base URL
     * @param {Object} payload - Data to send in the request body
     * @returns {Promise<any>} Response data from the server
     */
    async post(relativeUrl, payload) {
        const response = await axiosInstance.post(relativeUrl, payload);
        return response.data;
    },

    /**
     * Upload file(s) to the specified URL
     * @param {string} relativeUrl - The endpoint URL relative to base URL
     * @param {File|File[]} files - Single file or array of files to upload
     * @param {Function|null} onProgress - Optional callback for upload progress
     * @returns {Promise<any>} Response data from the server
     * 
     * @example
     * // Single file upload
     * const file = input.files[0];
     * await request.upload('/upload', file, (progress) => {
     *     console.log(`Upload progress: ${progress}%`);
     * });
     * 
     * // Multiple files upload
     * const files = Array.from(input.files);
     * await request.upload('/upload', files);
     */
    async upload(relativeUrl, formData, onProgress = null) {
        try {
            const response = await axiosInstance.post(relativeUrl, formData, {
                headers: {
                    'Content-Type': 'multipart/form-data'
                },
                onUploadProgress: onProgress ? (progressEvent) => {
                    const percentCompleted = Math.round((progressEvent.loaded * 100) / progressEvent.total);
                    onProgress(percentCompleted);
                } : undefined
            });
            return response.data;
        } catch (error) {
            console.error('Upload error:', error);
            throw new Error('Failed to upload file');
        }
    }
};

/**
 * Toast - global function to show toast notifications
 * @param {Object|string} options - Toast configuration or message string
 * @param {string} options.message - Message to display
 * @param {number} [options.duration=2000] - Duration in milliseconds
 * @param {string} [options.type='info'] - Type of toast (info|success|error|warning)
 * @param {string} [options.position='center'] - Position of toast (top|center|bottom)
 * 
 * @example
 * // Simple usage
 * showToast('Hello World');
 * 
 * // Advanced usage
 * showToast({
 *     message: 'Operation successful',
 *     type: 'success',
 *     duration: 3000,
 *     position: 'top'
 * });
 */
window.showToast = function(options) {
    const config = typeof options === 'string' 
        ? { message: options } 
        : options;

    const {
        message,
        duration = 2000,
        type = 'info',
        position = 'center'
    } = config;

    // Remove all existing toasts
    document.querySelectorAll('.toast-message').forEach(el => el.remove());

    // Create toast element
    const toast = document.createElement('div');
    toast.className = 'toast-message';  // Use class instead of id
    toast.innerHTML = message;

    // Position styles based on position parameter
    const positionStyles = {
        top: 'top: 20px; left: 50%; transform: translateX(-50%);',
        center: 'top: 50%; left: 50%; transform: translate(-50%, -50%);',
        bottom: 'bottom: 20px; left: 50%; transform: translateX(-50%);'
    };

    // Type styles
    const typeStyles = {
        info: 'background-color: rgba(0, 0, 0, 0.7);',
        success: 'background-color: rgba(40, 167, 69, 0.9);',
        error: 'background-color: rgba(220, 53, 69, 0.9);',
        warning: 'background-color: rgba(255, 193, 7, 0.9);'
    };

    toast.style.cssText = `
        position: fixed;
        ${positionStyles[position]}
        color: white;
        padding: 10px 20px;
        border-radius: 4px;
        z-index: 9999;
        ${typeStyles[type]}
        transition: opacity 0.3s ease-in-out;
    `;

    document.body.appendChild(toast);

    // Fade out animation
    setTimeout(() => {
        toast.style.opacity = '0';
        setTimeout(() => {
            toast?.parentNode?.removeChild(toast);
        }, 300);
    }, duration);
};

/**
 * Navigate to a specified path
 * - Removes leading slash from the path if present
 * - Constructs the final path using BASE_PATH if it exists
 * - Updates the window location to navigate
 * 
 * @param {string} path - The path to navigate to
 */
window.goto = function(path) {
     const BASE_PATH = '';
    // Remove leading slash if present
    const cleanPath = path.startsWith('/') ? path.slice(1) : path;
    
    // Construct the final path, only add BASE_PATH if it exists
    const finalPath = BASE_PATH ? `${BASE_PATH}/${cleanPath}` : `/${cleanPath}`;
    
    // Perform the navigation
    window.location.href = finalPath;
}; 